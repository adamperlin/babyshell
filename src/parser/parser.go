package parser
import (
  "scanner"
  "io"
  "errors"
  "os"
  "fmt"
  "bytes"
  "strings"
)
type BasicCommand struct {
  Args []string
}
type CommandList struct {
  Commands []*BasicCommand
  In string
  Out string
  Err string
  num_Commands int
  curr_cmd *BasicCommand
  Io_redirect_mode string
  Background bool
}
type Parser struct {
  s *scanner.Scanner
  buf struct {
    tok scanner.Token
    lit string
    n int
  }
}
func NewParser(r io.Reader) *Parser {
  return &Parser{s: scanner.NewScanner(r)}
}
func (p *Parser) scan()(tok scanner.Token, lit string) {
  if p.buf.n != 0 {
    p.buf.n = 0
    return p.buf.tok, p.buf.lit
  }
  tok,  lit = p.s.Scan()
  p.buf.tok, p.buf.lit = tok, lit
  return
}

func (p *Parser) unscan(){p.buf.n = 1}

func(p *Parser) scanIgnoreWhitespace()(tok scanner.Token, lit string){
  tok, lit = p.scan()
  if tok == scanner.WS {
    tok,lit = p.scan()
  }
  return
}

func(p *Parser) parseString()(error, string){
  var buff bytes.Buffer
  _, litquote := p.scanIgnoreWhitespace()
  //buff.WriteString(litquote)
  parsingString := true
  for parsingString {
    tok, lit := p.s.Scan()
    switch tok {
        case scanner.ESCAPEDLIT:
            char := strings.Replace(lit, "\\", "", -1)
              buff.WriteString(char)
        break
      case scanner.QUOTE:
        if (lit) == litquote {
          parsingString = false
          return nil, buff.String()
        }else {
          err := fmt.Errorf("Error: unescaped quote literal")
          return err, ""
        }
      default:
        buff.WriteString(lit)
        break
    }
  }
  ret := func(s string) string{
    return strings.Replace(s,"\n","", -1)
  }(buff.String())
  return nil, ret
}


func(p *Parser) Parse() (error, *CommandList) {
  var isScanningForArgs bool
  cl := &CommandList{}
  var currentcommand *BasicCommand = new(BasicCommand)
  cl.curr_cmd = currentcommand
  tok, lit := p.scanIgnoreWhitespace()
  if  tok != scanner.IDENT && tok != scanner.NEWLINE {
    err := fmt.Errorf("Error: expected command or pathname, found %q", lit)
    return err, nil
  }else if tok == scanner.NEWLINE {
    return nil,nil
  }
  currentcommand.Args = append(currentcommand.Args, lit)
  //scan for pipes, other io redirection
  for {
    tok, lit = p.scanIgnoreWhitespace()
    switch tok {
    case scanner.IDENT:
      currentcommand.Args = append(currentcommand.Args, lit)
      break
    case scanner.PIPE:
      tok, lit = p.scanIgnoreWhitespace()
      isScanningForArgs = true
      if tok != scanner.IDENT && tok != scanner.NEWLINE {
        fmt.Fprintf(os.Stderr, "Error: expected command to pipe to, found %q", lit)
        p.unscan()
        return errors.New("Unexpected Token"), nil
        break
      }
      cl.Commands = append(cl.Commands, currentcommand)
      currentcommand = new(BasicCommand)
      cl.curr_cmd = currentcommand
      currentcommand.Args = append(currentcommand.Args, lit)
      isScanningForArgs = false
      break
    case scanner.GREAT:
      tok,lit = p.scanIgnoreWhitespace()
      if tok == scanner.IDENT {
        cl.Out = lit
        break
      }else  {
        err := fmt.Errorf("Expected file name, got %q", lit)
        p.unscan()
        return err, nil
      }
    case scanner.LESS:
      tok,lit = p.scanIgnoreWhitespace()
      if tok == scanner.IDENT {
        cl.In = lit
      }else{
        err := fmt.Errorf("Error: expected pathname, found %q", lit)
        p.unscan()
        return err, nil
      }
    case scanner.NEWLINE:
      if isScanningForArgs {
        break
      }
      if cl.curr_cmd == currentcommand {
      cl.Commands = append(cl.Commands, currentcommand)
      }
      return nil, cl
    default:
      err := fmt.Errorf("Error: expected command or pathname, found %q", lit)
      return err, nil
    case scanner.AMPERSAND:
      tok, lit = p.scanIgnoreWhitespace()
      if tok != scanner.NEWLINE  {
        err := fmt.Errorf("Error: '&' must be placed at the end of a line")
        p.unscan()
        return err, nil
      }else {
        cl.Background = true
        p.unscan()
        break
      }
    case scanner.QUOTE:
      p.unscan()
      err, str := p.parseString()
      if err != nil {
        return err, nil
      }else {
        currentcommand.Args = append(currentcommand.Args, str)
      }
      break
    }
  }
}

	"strconv"
	unrecognized       parserState = "unrecognized"
	diffBegin          parserState = "diffBegin"
	diffBody           parserState = "diffBody"
	newFileDiff        parserState = "newFileDiff"
	deleteFileDiff     parserState = "deleteFileDiff"
	modifyFileDiff     parserState = "modifyFileDiff"
	modifyFileModeDiff parserState = "modifyFileModeDiff"
	ignoreDiff         parserState = "ignoreDiff"
		if strings.HasPrefix(line, "deleted file mode ") {
		if strings.HasPrefix(line, "new file mode ") {
		if strings.HasPrefix(line, "old mode ") {
			return p.handleModifyFileModeDiff(line)
		}
		return fmt.Errorf("unexpected diff line in state `%s`: %#v", p.state, line)
	case modifyFileModeDiff:
		if strings.HasPrefix(line, "new mode ") {
			p.state = unrecognized
			return p.writeOutLine(line)
		}
		return fmt.Errorf("unexpected diff line in state `%s`: %#v", p.state, line)

	a, b := lineParts[2], lineParts[3]

	trimmedPaths := make(map[string]string)

	for _, data := range []struct{ PathWithPrefix, Prefix string }{{a, "a/"}, {b, "b/"}} {
		if strings.HasPrefix(data.PathWithPrefix, "\"") && strings.HasSuffix(data.PathWithPrefix, "\"") {
			pathWithPrefix, err := strconv.Unquote(data.PathWithPrefix)
			if err != nil {
				return fmt.Errorf("unable to unqoute diff path %#v: %s", data.PathWithPrefix, err)
			}

			path := strings.TrimPrefix(pathWithPrefix, data.Prefix)
			if !p.PathFilter.IsFilePathValid(path) {
				p.state = ignoreDiff
				return nil
			}
			newPath := p.PathFilter.TrimFileBasePath(path)
			newPathWithPrefix := data.Prefix + newPath
			trimmedPaths[data.PathWithPrefix] = strconv.Quote(newPathWithPrefix)
		} else {
			path := strings.TrimPrefix(data.PathWithPrefix, data.Prefix)
			if !p.PathFilter.IsFilePathValid(path) {
				p.state = ignoreDiff
				return nil
			}
			newPath := p.PathFilter.TrimFileBasePath(path)
			trimmedPaths[data.PathWithPrefix] = data.Prefix + newPath
		}
	newLine := fmt.Sprintf("diff --git %s %s", trimmedPaths[a], trimmedPaths[b])
func (p *diffParser) handleModifyFileModeDiff(line string) error {
	p.state = modifyFileModeDiff
	return p.writeOutLine(line)
}

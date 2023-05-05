package main

import (
	"fmt"
	links "links_parser"
	"strings"
)

var testHtml = `
<!DOCTYPE html>
<html>
<head>
	<title>Pacifism and Buddhism</title>
</head>
<body>
	<h1>Pacifism and Buddhism</h1>
	<ul>
		<li><a href="https://en.wikipedia.org/wiki/Pacifism">Pacifism (Wikipedia)</a></li>
		<li><a href="https://www.buddhanet.net/e-learning/history/schools.htm">Buddhist Schools and Sects (BuddhaNet)</a></li>
		<li><a href="https://plato.stanford.edu/entries/pacifism/">Pacifism (Stanford Encyclopedia of Philosophy)</a></li>
		<li><a href="https://www.buddhanet.net/e-learning/5minbud.htm">What is Buddhism? (BuddhaNet)</a></li>
		<li><a href="https://www.peaceparks.org/">Peace Parks Foundation</a></li>
		<li><a href="https://www.buddhanet.net/e-learning/buddhistworld/index.htm">Buddhism Around the World (BuddhaNet)</a></li>
		<li><a href="https://www.nonviolence.org/">International Center on Nonviolent Conflict</a></li>
		<li><a href="https://www.buddhanet.net/e-learning/history/briefhis.htm">A Brief History of Buddhism (BuddhaNet)</a></li>
		<li><a href="https://www.icrc.org/en">International Committee of the Red Cross</a></li>
		<li><a href="https://www.buddhanet.net/e-learning/basic-guide.htm">A Basic Guide to Meditation (BuddhaNet)</a></li>
		<li><a href="https://www.wri-irg.org/en">War Resisters' International</a></li>
		<li><a href="https://www.buddhanet.net/audio-library.htm">Buddhist Audio Library (BuddhaNet)</a></li>
		<li><a href="https://www.sgi.org/">Soka Gakkai International</a></li>
		<li><a href="https://www.buddhanet.net/e-learning/8foldpath.htm">The Noble Eightfold Path (BuddhaNet)</a></li>
		<li><a href="https://www.ipb.org/">International Peace Bureau</a></li>
		<li><a href="https://www.buddhanet.net/e-learning/buddhistdictionary.htm">Buddhist Dictionary (BuddhaNet)</a></li>
		<li><a href="https://www.icj.org/">International Court of Justice</a></li>
	</ul>
</body>
</html>

`

func main() {
	r := strings.NewReader(testHtml)
	links, err := links.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}

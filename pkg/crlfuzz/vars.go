package crlfuzz

var (
	appendList = [...]string{
		"",
		"crlfuzz",
		"?crlfuzz=",
		// We use the fragment in another project to differentiate other URLs, so we omit this one here
		// "#",
	}

	escapeList = [...]string{
		"%00",
		"%0a",
		"%0a%20",
		"%0d",
		"%0d%09",
		"%0d%0a",
		"%0d%0a%09",
		"%0d%0a%20",
		"%0d%20",
		"%20",
		"%20%0a",
		"%20%0d",
		"%20%0d%0a",
		"%23%0a",
		"%23%0a%20",
		"%23%0d",
		"%23%0d%0a",
		"%23%0a",
		"%25%30",
		"%25%30%61",
		"%2e%2e%2f%0d%0a",
		"%2f%2e%2e%0d%0a",
		"%2f..%0d%0a",
		"%3f",
		"%3f%0a",
		"%3f%0d",
		"%3f%0d%0a",
		"%e5%98%8a%e5%98%8d",
		"%e5%98%8a%e5%98%8d%0a",
		"%e5%98%8a%e5%98%8d%0d",
		"%e5%98%8a%e5%98%8d%0d%0a",
		"%e5%98%8a%e5%98%8d%e5%98%8a%e5%98%8d",
		"%u0100",
		"%u010a",
		"%u010d",
		"%u010d%u010a",
	}
)

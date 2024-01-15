package templates

var HeaderAll = HeaderProps{
	Brand: HeaderBrand{
		Link: "/",
		Icon: IconTux,
		Text: "Linux World",
	},
	Elements: []NavbarElement{
		{
			Link:   "#",
			Text:   "Docs",
			Active: false,
		},
		{
			Link:   "#",
			Text:   "Blog",
			Active: true,
		},
		{
			Link:   "#",
			Text:   "Contact us",
			Active: false,
		},
	},
	Links: []NavbarSocialLink{
		{
			Link:  "#",
			Icon:  IconGitHub,
			Title: "GitHub",
		},
		{
			Link:  "#",
			Icon:  IconDiscord,
			Title: "Discord",
		},
		{
			Link:  "#",
			Icon:  IconYouTube,
			Title: "YouTube",
		},
	},
}

var SidebarUser = SidebarProps{
	Elements: []SidebarElement{
		{
			Icon: IconCat,
			Text: "Cats",
			Link: "/cats",
		},
		{
			Icon: IconAlien,
			Text: "Aliens",
			Link: "/aliens",
		},
		{
			Icon: IconMoon,
			Text: "Space",
			Link: "/space",
		},
		{
			Icon: IconShuttle,
			Text: "Shuttle",
			Link: "/shuttle",
		},
		{
			Icon: IconStars,
			Text: "Theme",
			Link: "#",
		},
	},
}

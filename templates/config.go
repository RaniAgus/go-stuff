package templates

var HeaderAll = HeaderProps{
	Brand: HeaderBrand{
		Link: "/",
		Icon: IconTux,
		Text: "Linux World",
	},
	Elements: []NavbarElement{
		{
			Link: "#",
			Text: "Docs",
		},
		{
			Link: "#",
			Text: "Blog",
		},
		{
			Link: "#",
			Text: "Contact us",
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
			Link: "#",
		},
		{
			Icon: IconAlien,
			Text: "Aliens",
			Link: "#",
		},
		{
			Icon: IconMoon,
			Text: "Space",
			Link: "#",
		},
		{
			Icon: IconShuttle,
			Text: "Shuttle",
			Link: "#",
		},
		{
			Icon: IconStars,
			Text: "Theme",
			Link: "#",
		},
	},
}

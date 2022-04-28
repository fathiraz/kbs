// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
	title: 'Kitabisa test documentation',
	tagline: 'Kitabisa test',
	url: 'https://fathiraz.github.io',
	baseUrl: '/kbs/',
	onBrokenLinks: 'throw',
	onBrokenMarkdownLinks: 'warn',
	favicon: 'img/favicon_ktbs.ico',
	organizationName: 'fathiraz', // Usually your GitHub org/user name.
	projectName: 'kbs', // Usually your repo name.
	presets: [
		[
			'classic',
			/** @type {import('@docusaurus/preset-classic').Options} */
			({
				docs: {
					sidebarPath: require.resolve('./sidebars.js'),
					// Please change this to your repo.
					editUrl: 'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
				},
				blog: {
					showReadingTime: true,
					// Please change this to your repo.
					editUrl:
						'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
				},
				theme: {
					customCss: require.resolve('./src/css/custom.css'),
				},
			}),
		],
	],

	themeConfig:
	/** @type {import('@docusaurus/preset-classic').ThemeConfig} */
		({
			navbar: {
				title: 'Kitabisa',
				logo: {
					alt: 'Kitabisa logo',
					src: 'img/logogram__ktbs_white.png',
				},
				items: [
					{
						type: 'doc',
						docId: 'problem-3/question',
						position: 'left',
						label: 'Documentation',
					},
					{
						href: 'https://github.com/facebook/docusaurus',
						label: 'GitHub',
						position: 'right',
					},
				],
			},
			footer: {
				style: 'dark',
				links: [
					{
						title: 'Test',
						items: [
							{
								label: 'Link',
								href: 'https://drive.google.com/drive/folders/1BbssemnrRGOFd4rkE14r1DQTiGC0MN9l',
							},
						],
					},
					{
						title: 'Github',
						items: [
							{
								label: 'Link',
								href: 'https://github.com/fathiraz/kbs',
							},
						],
					},
					{
						title: 'Site',
						items: [
							{
								label: 'Kitabisa',
								href: 'https://kitabisa.com',
							},
						],
					},
				],
				copyright: `Copyright © 2022 test@kitabisa. Built with Docusaurus.`,
			},
			prism: {
				theme: lightCodeTheme,
				darkTheme: darkCodeTheme,
			},
		}),
};

module.exports = config;

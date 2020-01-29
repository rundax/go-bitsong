module.exports = {
    title: "BitSong Network",
    description: "Welcome to the BitSong developer site!",
    base: process.env.VUEPRESS_BASE || "/",
    themeConfig: {
        nav: [
            { text: 'Website', link: 'https://bitsong.io' },
            { text: 'Player Demo', link: 'https://demo.bitsong.io' },
            { text: 'Testnet Explorer', link: 'https://testnet.explorebitsong.com' },
            { text: 'Community', link: 'https://btsg.community' },
            { text: 'Blog', link: 'https://medium.com/@bitsongofficial' },
            { text: 'Github', link: 'https://github.com/bitsongofficial' },
        ],
        sidebar: [
            {
                title: "Guide",
                path: "/guide/",
                children: [
                    ['/guide/installation.md', 'Install BitSong Network'],
                    ['/guide/join-testnet.md', 'Join Public Testnet'],
                    ['/guide/upgrade-node.md', 'Upgrade Your Node'],
                    ['/guide/bitsongcli.md', 'BitSong CLI'],
                ]
            },
            {
                title: "Delegators",
                path: "/delegators/",
                children: [
                    ['/delegators/delegator-guide-cli.md', 'Delegator Guide (CLI)'],
                    ['/delegators/delegator-faq.md', 'Delegator FAQ'],
                    ['/delegators/delegator-security.md', 'Delegator Security']
                ]
            },
            {
                title: "Validators",
                path: "/validators/",
                children: [
                    ['/validators/overview.md', 'Overview'],
                    ['/validators/validator-setup.md', 'Setting Up a Validator'],
                    ['/validators/validator-faq.md', 'Validator FAQ'],
                    ['/validators/validator-security.md', 'Validator Security'],
                ]
            }]
    },
    plugins: [
        [
            "@vuepress/google-analytics",
            {
                ga: "UA-11653102-15"
            }
        ],
        [
            "sitemap",
            {
                hostname: "https://btsg.dev"
            }
        ]
    ]
};
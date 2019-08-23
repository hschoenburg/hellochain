module.exports = {
  title: "HelloChain",
  description: "A Hello World tutorial for building blockchains with Cosmos",
  ga: "",
  dest: "./dist",
  base: "/",
  markdown: {
    lineNumbers: true,
  },
  themeConfig: {
    repo: "cosmos/hellochain",
    editLinks: true,
    docsDir: "",
    docsBranch: "master",
    editLinkText: 'Edit this page on Github',
    lastUpdated: true,
    nav: [
      { text: "Back to Cosmos", link: "https://cosmos.network" },
      { text: "SDK Documentation", link: "https://cosmos.network/docs" }
    ],
    sidebar: [
      {
        title: "Simple App",
        collapsable: false,
        children: [
          "/tutorial/00-intro.md",
          "/tutorial/01-simple-app.md",
          "/tutorial/02-simple-start.md",
          "/tutorial/03-make.md",
          "/tutorial/04-try-it-out.md",
        ]
      }, {
      title: "Custom Module",
      collapsable: false,
      children: [
          "/tutorial/10-types.md",
          "/tutorial/11-module.md",
          "/tutorial/12-msgs.md",
          "/tutorial/13-handler.md",
          "/tutorial/14-keeper.md",
          "/tutorial/15-querier.md",
          "/tutorial/16-client.md",
          "/tutorial/17-cli.md",
          "/tutorial/18-full-cmd.md",
          "/tutorial/19-full-app.md",
          "/tutorial/20-complete.md",

        ]
      },
    ]
  }
}

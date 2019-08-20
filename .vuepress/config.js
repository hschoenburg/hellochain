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
          "/tutorial/intro.md",
          "/tutorial/simple-app.md",
          "/tutorial/simple-start.md",
          "/tutorial/make.md",
          "/tutorial/try-it-out.md",
        ]
      }, {
      title: "Custom Module",
      collapsable: false,
      children: [
          "tutorial/types.md",
          "/tutorial/module.md",
          "/tutorial/msgs.md",
          "/tutorial/handler.md",
          "/tutorial/keeper.md",
          "/tutorial/querier.md",
        "/tutorial/client.md",
        "/tutorial/cli.md",
          "/tutorial/full-app.md",
          "/tutorial/full-cmd.md",
          "/tutorial/complete.md",

        ]
      },
    ]
  }
}

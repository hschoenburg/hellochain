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
    docsDir: "whatisthis",
    docsBranch: "develop",
    editLinkText: 'Edit this page on Github',
    lastUpdated: true,
    nav: [
      { text: "Back to Cosmos", link: "https://cosmos.network" },
      { text: "SDK Documentation", link: "https://cosmos.network/docs" }
    ],
    sidebar: [
      {
        title: "Hellochain Tutorial",
        collapsable: false,
        children: [
          "/tutorial/basic-app.md",
          "/tutorial/simple-start.md",
          "/tutorial/make.md",
          "/tutorial/msgs.md",
          "/tutorial/handler.md",
        ]
      },
    ]
  }
}

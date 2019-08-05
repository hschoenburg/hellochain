module.exports = {
  title: "Cosmos Tutorials",
  description: "A collection of tutorials for building blockchains with the Cosmos SDK.",
  ga: "",
  dest: "./dist",
  base: "/",
  markdown: {
    lineNumbers: true
  },
  themeConfig: {
    repo: "cosmos/hellochain",
    editLinks: true,
    docsDir: "whatisthis",
    docsBranch: "develop",
    editLinkText: 'Edit this page on Github',
    lastUpdated: true,
    algolia: {
      apiKey: 'a6e2f64347bb826b732e118c1366819a',
      indexName: 'cosmos_network',
      debug: false
    },
    nav: [
      { text: "Back to Cosmos", link: "https://cosmos.network" },
      { text: "SDK Documentation", link: "https://cosmos.network/docs" }
    ],
    sidebar: [
      {
        title: "Hellochain Tutorial",
        collapsable: true,
        children: [
          "/tutorial/",
          "/tutorial/app.md",
          "/tutorial/keeper.md",
          "/tutorial/querier.md",
          "/tutorial/handler.md",
        ]
      },
      {
        title: "Nameservice Tutorial",
        collapsable: true,
        children: [
          "/nameservice/",
          "/nameservice/app-complete.md",
          "/nameservice/codec",
        ]
      },
    ]
  }
}

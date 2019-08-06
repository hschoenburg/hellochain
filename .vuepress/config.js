module.exports = {
  title: "HelloChain",
  description: "A Hello World tutorial for building blockchains with Cosmos",
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
        collapsable: false,
        children: [
          "/tutorial/",
          "/tutorial/app.md",
          "/tutorial/handler.md",
        ]
      },
    ]
  }
}

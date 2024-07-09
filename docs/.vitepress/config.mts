import { defineConfig } from "vitepress";
import { tabsMarkdownPlugin } from "vitepress-plugin-tabs";

// @ts-ignore
const gaTagId = process.env.GA_TAG_ID;

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "OutageLab",
  description: "OutageLab",
  appearance: "dark",
  markdown: {
    config(md) {
      md.use((mdit) => {
        //@ts-ignore
        tabsMarkdownPlugin(mdit);
        return md;
      });
    },
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "Docs", link: "/docs/intro/" },
      {
        component: "NavButton",
        props: {
          text: "Go to App",
          theme: "brand",
          href: "https://app.outagelab.com",
          rel: "sponsored",
        },
      },
    ],

    sidebar: [
      {
        text: "Introduction",
        items: [
          {
            text: "What is OutageLab?",
            link: "/docs/intro/",
          },
          {
            text: "Quickstart",
            link: "/docs/intro/quickstart/",
          },
          {
            text: "Usecases",
            link: "/docs/intro/usecases",
          },
        ],
      },
      {
        text: "Features",
        items: [{ text: "Outage Types", link: "/docs/features/outage-types" }],
      },
      {
        text: "Reference",
        items: [
          {
            text: "Integration SDK's",
            link: "/docs/reference/integration-sdks",
            items: [
              {
                text: "Go",
                link: "/docs/reference/integration-sdks/go",
              },
              {
                text: "Python",
                link: "/docs/reference/integration-sdks/python",
              },
              {
                text: "Node.js",
                link: "/docs/reference/integration-sdks/javascript",
              },
              {
                text: "Other Languages",
                link: "/docs/reference/integration-sdks/other",
              },
              // { text: "Golang", link: "/golang-sdk" },
              // { text: "Python", link: "/python-sdk" },
              // { text: "Ruby", link: "/ruby-sdk" },
              // { text: "Java", link: "/java-sdk" },
            ],
          },
        ],
      },
    ],

    socialLinks: [
      { icon: "github", link: "https://github.com/outagelab/outagelab" },
    ],
    search: {
      provider: "local",
    },
  },
  head: gaTagId
    ? [
        [
          "script",
          {
            async: "",
            src: `https://www.googletagmanager.com/gtag/js?id=${gaTagId}`,
          },
        ],
        [
          "script",
          {},
          `window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());
      gtag('config', '${gaTagId}');`,
        ],
      ]
    : [],
});

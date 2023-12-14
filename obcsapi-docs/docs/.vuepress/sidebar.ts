import { sidebar } from "vuepress-theme-hope";

export default sidebar({
  "/": [
    "",
    "md",
    {
      text: "Go 版本",
      icon: "book",
      prefix: "md/go-version/",
      children: "structure",
      link: "md/go-version",
    },
    {
      text: "Python 版本",
      icon: "book",
      link: "md/python-version",
    },
    "md/faq"
  ],
});

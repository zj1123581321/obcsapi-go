if(!self.define){let e,s={};const i=(i,n)=>(i=new URL(i+".js",n).href,s[i]||new Promise((s=>{if("document"in self){const e=document.createElement("script");e.src=i,e.onload=s,document.head.appendChild(e)}else e=i,importScripts(i),s()})).then((()=>{let e=s[i];if(!e)throw new Error(`Module ${i} didn’t register its module`);return e})));self.define=(n,r)=>{const l=e||("document"in self?document.currentScript.src:"")||location.href;if(s[l])return;let t={};const o=e=>i(e,l),u={module:{uri:l},exports:t,require:o};s[l]=Promise.all(n.map((e=>u[e]||o(e)))).then((e=>(r(...e),t)))}}define(["./workbox-3625d7b0"],(function(e){"use strict";self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"assets/AboutView-4bd8d9e2.js",revision:null},{url:"assets/AboutView-4d995ba2.css",revision:null},{url:"assets/index-d63dd564.js",revision:null},{url:"assets/index-f9c63b4c.css",revision:null},{url:"assets/LoginView-802caac8.js",revision:null},{url:"assets/Setting-8e0e8056.js",revision:null},{url:"assets/TalkView-670796a4.css",revision:null},{url:"assets/TalkView-c6a0584b.js",revision:null},{url:"index.html",revision:"530ea144c28e81862338c76f1d3f9ead"},{url:"registerSW.js",revision:"bf6c2f29aef95e09b1f72cf59f427a55"},{url:"style.css",revision:"2d898f5a9d916da9106b2c0e57e01076"},{url:"pwa-192x192.png",revision:"72a628e05914d3a0436e811db691b653"},{url:"manifest.webmanifest",revision:"af9749ee2a07a33170f70a1416ce6de9"}],{}),e.cleanupOutdatedCaches(),e.registerRoute(new e.NavigationRoute(e.createHandlerBoundToURL("index.html")))}));

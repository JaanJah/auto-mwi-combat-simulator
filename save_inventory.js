// ==UserScript==
// @name         MWI - Save Inventory
// @namespace    http://tampermonkey.net/
// @version      0.0.2
// @description  Get MWI Inventory and save as JSON
// @author       JaanJah
// @icon         https://www.google.com/s2/favicons?sz=64&domain=milkywayidle.com
// @match        https://www.milkywayidle.com/game?characterId=*
// @match        https://test.milkywayidle.com/game?characterId=*
// @grant        none
// ==/UserScript==

(function () {
  "use strict";

  hookWS();
  waitForElement(".NavigationBar_minorNavigationLinks__dbxh7", (element) => {
    let copyButton = createCopyButton(element);
    copyButton.addEventListener("click", () => copyToClipboard(copyButton));
  });
  let client_data = {};
  let character_data = {};

  // WebSocket hook logic taken from MooneyCalc Importer
  // @link https://greasyfork.org/en/scripts/494468-mooneycalc-importer
  function hookWS() {
    const dataProperty = Object.getOwnPropertyDescriptor(MessageEvent.prototype, "data");
    const oriGet = dataProperty.get;

    dataProperty.get = hookedGet;
    Object.defineProperty(MessageEvent.prototype, "data", dataProperty);

    function hookedGet() {
      const socket = this.currentTarget;
      if (!(socket instanceof WebSocket)) {
        return oriGet.call(this);
      }
      if (socket.url.indexOf("api.milkywayidle.com/ws") <= -1
        && socket.url.indexOf("api-test.milkywayidle.com/ws") <= -1) {
        return oriGet.call(this);
      }

      const message = oriGet.call(this);
      Object.defineProperty(this, "data", { value: message }); // Anti-loop

      return handleMessage(message);
    }
  }

  function handleMessage(message) {
    let obj = JSON.parse(message);
    if (obj?.type === "init_character_data") {
      character_data = obj;
    }
    // Useful for creating mapping in simulator
    // if (obj?.type === "init_client_data") {
    //   client_data = obj;
    // }
    return message;
  }

  function waitForElement(selector, callback, intervalTime = 500, timeout = 10000) {
      const startTime = Date.now();

      const interval = setInterval(() => {
          const element = document.querySelector(selector);
          if (element) {
              clearInterval(interval);
              callback(element);
          } else if (Date.now() - startTime > timeout) {
              clearInterval(interval);
              console.warn(`Element ${selector} not found within ${timeout}ms`);
          }
      }, intervalTime);
  }

  function createCopyButton(parent) {
    const childClass = "NavigationBar_minorNavigationLink__31K7Y";
    const copyButton = document.createElement("div")
    copyButton.classList.add(childClass)
    copyButton.innerHTML = "Copy Inventory to Clipboard"
    parent.appendChild(copyButton)
    return copyButton;
  }

  function copyToClipboard(copyButton) {
    navigator.clipboard.writeText(JSON.stringify({
      // client_data,
      character_data
    }));
    const originalText = copyButton.innerHTML;
    copyButton.innerHTML = "Copied!";
    setTimeout(() => {
      copyButton.innerHTML = originalText;
    }, 500);
  }
})();

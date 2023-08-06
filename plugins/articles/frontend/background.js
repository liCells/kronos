import { Readability } from './Readability.js';

chrome.action.onClicked.addListener(tab => {
  chrome.scripting.executeScript({
    target: { tabId: tab.id },
    files: ['content.js'],
  }, () => {
    chrome.tabs.executeScript(tab.id, {
      code: `window.Readability = ${Readability.toString()};`
    });
  });
});


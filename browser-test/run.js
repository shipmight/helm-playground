const puppeteer = require("puppeteer");
const httpServer = require("http-server");
const path = require("path");
const assert = require("assert");

const serverHost = "127.0.0.1";
const serverPort = 8123;
const serverUrl = `http://${serverHost}:${serverPort}`;

const initialOutput = "---\nexample:\n  - first\n  - second\n";
const initialHash =
  "t=LQhQFMA8EMFsAcA24BcACA3h4aB0A1aRAV3AGdcBLAF3FjLQB81qB7ATTkSbQDtLeAE3C9qaAExoAvlNBA&v=LQhQEsBcFMFsGcBcACA2gIgGbgE70ugDTLrzQDGA9gHYAm6AuqEA";
const changeValuesContentTo = 'items: ["hello", "from", "puppeteer"]';
const expectedOutput = "---\nexample:\n  - hello\n  - from\n  - puppeteer\n";
const expectedHash =
  "t=LQhQFMA8EMFsAcA24BcACA3h4aB0A1aRAV3AGdcBLAF3FjLQB81qB7ATTkSbQDtLeAE3C9qaAExoAvlNBA&v=JYFwpgtgzgXABAbQEQAswBt0HskBo5IBmATlhHgQA4CullY4YxSAukA";

async function run() {
  console.error("[browser-test] starting server");
  const server = httpServer.createServer({
    root: path.join(__dirname, ".."),
    cache: -1,
  });
  await new Promise((resolve, reject) => {
    server.listen(serverPort, serverHost, (err) => {
      if (err) {
        return reject(err);
      }
      resolve();
    });
  });
  console.error(`[browser-test] server started at ${serverUrl}`);

  console.error("[browser-test] starting browser");
  const browser = await puppeteer.launch({
      args: ['--no-sandbox', '--disable-setuid-sandbox']
  });
  const page = await browser.newPage();
  await page.goto(serverUrl);
  console.error("[browser-test] started browser");

  const getPageUrl = async () =>
    page.evaluate(() => window.location.toString());
  const setValuesContent = async (value) =>
    page.$eval(
      ".textarea--values",
      (el, value) => el.CodeMirror.setValue(value),
      value
    );
  const getOutput = async () =>
    page.$eval(".textarea--output", (el) => el.CodeMirror.getValue());

  console.error("[browser-test] running test");
  await page.click(".start");
  await page.waitForSelector(".start", { hidden: true, timeout: 5000 });
  assert.equal(await getOutput(), initialOutput, "should see initial output");
  assert.equal(
    await getPageUrl(),
    `${serverUrl}/#${initialHash}`,
    "should see initial URL hash"
  );
  await setValuesContent(changeValuesContentTo);
  assert.equal(await getOutput(), expectedOutput, "should see updated output");
  assert.equal(
    await getPageUrl(),
    `${serverUrl}/#${expectedHash}`,
    "should see updated URL hash"
  );
  console.error("[browser-test] ran test");

  console.error("[browser-test] closing browser");
  await browser.close();
  console.error("[browser-test] closed browser");

  console.log("[browser-test] PASS");
  process.exit(0);
}

run().catch((error) => {
  console.error("[browser-test] ERROR ", error);
  process.exit(1);
});

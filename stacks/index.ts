import { App } from "@serverless-stack/resources";
import { MyStack } from "./MyStack";

export default function (app: App) {
  app.setDefaultFunctionProps({
    runtime: "go1.x",
    srcPath: "api",
    environment: { MONGODB_URI: process.env.MONGODB_URI || "" },
  });
  app.stack(MyStack);
}

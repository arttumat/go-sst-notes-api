import { Api, StackContext } from "@serverless-stack/resources";

export function MyStack({ stack }: StackContext) {
  // Create the HTTP API
  const api = new Api(stack, "api", {
    routes: {
      "GET /notes": "functions/list.go",
      "POST /notes": "functions/create.go",
      "GET /notes/{id}": "functions/get.go",
      "PUT /notes/{id}": "functions/update.go",
    },
  });

  // Show API endpoint in output
  stack.addOutputs({
    ApiEndpoint: api.url,
  });
}

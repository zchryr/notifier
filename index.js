const core = require('@actions/core'); // Importing core module from GitHub Actions toolkit
const axios = require('axios'); // Importing axios for making HTTP requests
const github = require('@actions/github'); // Importing GitHub toolkit for context information

async function run() {
  try {
    // Reading input parameters from the workflow file
    const userInput = core.getInput('input'); // User-provided JSON input
    const url = core.getInput('url'); // URL to which data is sent
    const expectedResponseCode = core.getInput('response_code'); // Expected HTTP response code

    // Validate user input is valid JSON
    let inputJson;
    try {
      inputJson = JSON.parse(userInput); // Attempt to parse user input as JSON
    } catch (error) {
      throw new Error('Invalid JSON input'); // Throw error if JSON parsing fails
    }

    // Constructing the git_info object with workflow and repo info
    const gitInfo = {
      workflow: github.context.workflow, // Fetching the name of the current workflow
      repo: github.context.repo.repo, // Fetching the name of the repository
    };

    // Constructing the body of the data to be sent
    const body = {
      git_info: gitInfo, // Including git_info in the body
      input: inputJson, // Including parsed user input in the body
    };

    // Sending data to the specified URL using a POST request
    const response = await axios.post(url, body);

    // Checking if the actual response code matches the expected response code
    if (response.status.toString() !== expectedResponseCode) {
      // Throwing an error if the response codes don't match
      throw new Error(
        `Expected response code ${expectedResponseCode}, but got ${response.status}`
      );
    }

    // Setting the 'success' output parameter to true if all goes well
    core.setOutput('success', true);
  } catch (error) {
    // If any error occurs, mark the action as failed and log the error message
    core.setFailed(`Action failed with error: ${error.message}`);
  }
}

run();

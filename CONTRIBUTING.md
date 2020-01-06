# Contribution Guidelines

Contributions to this Module are very welcome! We follow a fairly standard pull request process for contributions, subject to the following guidelines:

1. [Update the tests](#update-the-tests)
1. [Update the code](#update-the-code)
1. [Create a pull request](#create-a-pull-request)
1. [Merge and release](#merge-and-release)

## Update the tests

We also recommend updating the automated tests before updating any code (see Test Driven Development). That means you add or update a test case, verify that it's failing with a clear error message, and then make the code changes to get that test to pass. This ensures the tests stay up to date and verify all the functionality in this Module, including whatever new functionality you're adding in your contribution. Check out the tests folder for instructions on running the automated tests.

## Update the Code

At this point, make your code changes and use your new test case to verify that everything is working. As you work, keep in mind about Backward Compatibility:

#### Backwards compatibility

Please make every effort to avoid unnecessary backwards incompatible changes. With Terraform code, this means:

1. Do not delete, rename, or change the type of input variables.
1. If you add an input variable, it should have a `default`.
1. Do not delete, rename, or change the type of output variables.
1. Do not delete or rename a module in the `modules` folder.

If a backwards incompatible change cannot be avoided, please make sure to call that out when you submit a pull request,
explaining why the change is absolutely necessary.


## Create a Pull Request

[Create a pull request](https://help.github.com/articles/creating-a-pull-request/) with your changes. Please make sure
to include the following:

1. A description of the change, including a link to your GitHub issue.
1. The output of your automated test run, preferably in a [GitHub Gist](https://gist.github.com/). We cannot run
   automated tests for pull requests automatically due to [security
   concerns](https://circleci.com/docs/fork-pr-builds/#security-implications), so we need you to manually provide this
   test output so we can verify that everything is working.
1. Any notes on backwards incompatibility or downtime.


## Merge and release

The maintainers for this repo will review your code and provide feedback. If everything looks good, they will merge the
code and release a new version, which you'll be able to find in the [releases page](../../releases).

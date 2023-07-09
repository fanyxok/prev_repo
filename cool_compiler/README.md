# Get Started

To get started, please simply fork this GitLab repository and
follow the structure, testing and submissions guidelines below. 
**Remember to change it to a private repo.**

# Repository Structure

### Define your dependencies

Define your dependencies in a apt.txt file.

Please make sure to always have a reference to `git` and `flex`

### Define a compiler.json

Define the following properties in the compiler.json:

- assignment_id: "project-assignment-1-2018fa" for PA1
- authors: your school mail address
- description: a description of your submission
  (optional)


# ~~Test~~ (Still under development)


Make sure you have installed `ShanghaiTech-compiler-judger` first by running `python3 -m pip install ShanghaiTech-compiler-judger`, this package is only tested on ubuntu and MacOS, if you are using Windows, maybe you need to have a look at the source code `http://s3l.shanghaitech.edu.cn:8081/compiler/compiler-client`

~~Within the root-folder of this repository, simple run:`compiler-test .`~~


# Submit

Submissions are done by simply running:

```compiler-submit /path/to/your/repo```

or simply `compiler-submit .` from within the root folder of this repo.

You can define your own submisson tag via `--tag your_tag`, 
otherwise a random one will be generated.

If you define your own, please use a new tag for every new submission.
Every submission will create a new GitLab issue, where you can track the progress.
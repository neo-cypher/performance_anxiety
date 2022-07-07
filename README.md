# Performance Anxiety

Performance Anxiety is a tool created to help you measure and track the performance of your codebase. In an effort to verify that your project is scalable and so that you can determine which features require optimization, we have created a tool that will help you measure performance. 

In furture releases, we will be adding the ability to track your project over time so that you can see when and which features are over-utilizing resources. Being able to compare your project performance over time will help you determine which features need to be optimized. Additionally, it will help you to determine when and where the code change that impacted resources was merged.

## Getting Started

### Step 1: Install the Tool

``` 
go install github.com/storj-antonio/performance-anxiety
```

### Step 2: Run the Tool

```
performance-anxiety -h hostname -e export_format
```

### Step 3: Results

```
performance-anxiety --results filename.json --export-format json
```

## Version History

* 0.1
    * Initial Release

## Acknowledgements

Inspiration, code snippets, etc.
* [awesome-readme](https://github.com/matiassingers/awesome-readme)
* [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)
* [dbader](https://github.com/dbader/readme-template)

## License

This project is licensed under the [NAME HERE] License - see the LICENSE.md file for details



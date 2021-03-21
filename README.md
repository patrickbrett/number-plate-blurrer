# Number Plate Blurrer

AI that detects and automagically blurs number plates in car photos.

Use it before posting photos of your car on social media, sales listings, etc.

Inspired by [Carsales' _Mystique_ AI](https://www.itnews.com.au/news/carsales-uses-ai-to-crack-down-on-number-plate-cloning-529790), but this one's open source!

## Motivations

Made because I needed to blur some photos when selling my car - why do it manually in Photoshop in 20 minutes when you can build an AI to do it for you in 20 hours? 🙌🏻

Beyond that, I wanted to challenge myself to build an AI that can learn from its mistakes and automatically re-train itself. To do this, every photo uploaded to the site provides the opportunity for the user to correct mistakes and improve the AI in the process.

## See it live

[not just yet - still got a bit of work to do!]

As of 21 March 2021 it's around 40% done.

## Built with

This was a learning project - my first exposure to Angular on the front end (as opposed to React) and Go on the back end (as opposed to Node/Java).

I've done everything I can to make the code as clean and modular as possible, but feedback is welcome!

Powered by: AWS, Lambda, SageMaker, Jupyter, Golang, Python, Angular.

You can host this all yourself on AWS, by following the instructions below.

## Getting it running yourself

All infrastructure is also in this repo in the form of CloudFormation templates. The goal is to have a single one-click template to deploy from, but for now it's held in a variety of templates.

### Commands

Bit of a mish mash here, I will clean this up soon.

```
aws s3 cp --recursive ./ s3://npb-source-bucket-sourcebucket-96mim45irkcp
```

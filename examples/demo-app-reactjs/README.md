# Deploying the app
1. Update dependencies
    ```
    $ npm install
    ```
1. Build
    ```
    $ npm run-script build
    ```
1. Push
    ```
    $ cf push
    ```

## Notes
- You can use the nodeJS buildpack or static file buildpack to deploy this app
- Static file buildpack is used in the example because it uses less resources at runtime (~8MB RAM for staticfile vs 300+MB RAM for nodejs)

## Useful Links
- Create React App: https://github.com/facebook/create-react-app
- NodeJS buildpack: https://docs.cloudfoundry.org/buildpacks/node/index.html
- Static file buildpack: https://docs.cloudfoundry.org/buildpacks/staticfile/index.html

---

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).
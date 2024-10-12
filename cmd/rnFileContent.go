package cmd

var TWConfig = `// tailwind.config.js

module.exports = {
  content: ["./App.{js,jsx,ts,tsx}", "./**/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {},
  },
  plugins: [],
}`

var Eslint = `module.exports = {
  root: true,
  extends: '@react-native',
  parserOptions: {
    requireConfigFile: false,
  },
};
`

var TWBabel = `// babel.config.js
module.exports = {
  presets: ['module:metro-react-native-babel-preset'],
  plugins: ["nativewind/babel"],
};`

var Nenvdts = `/// <reference types="nativewind/types" />`

var BabelFull = `module.exports = {
  presets: ['module:@react-native/babel-preset'],
  plugins: [
    'nativewind/babel',
    [
        "module:react-native-dotenv", {
        "envName": "APP_ENV",
        "moduleName": "@env",
        "path": ".env",
        "safe": false,
        "allowUndefined": true,
        "verbose": false
      }
    ],
    'react-native-reanimated/plugin'],
};
`
var envdts = `declare module '@env' {
    export const BASE_URL: string;
    
 }`

var esConfig = `// eslint.config.js
// export default [
//     {
//         // ignores: [".config/"]
//     }
// ];`

var buildGradel = `apply plugin: "com.android.application"
apply plugin: "org.jetbrains.kotlin.android"
apply plugin: "com.facebook.react"

/**
 * This is the configuration block to customize your React Native Android app.
 * By default you don't need to apply any configuration, just uncomment the lines you need.
 */
react {
    /* Folders */
    //   The root of your project, i.e. where "package.json" lives. Default is '../..'
    // root = file("../../")
    //   The folder where the react-native NPM package is. Default is ../../node_modules/react-native
    // reactNativeDir = file("../../node_modules/react-native")
    //   The folder where the react-native Codegen package is. Default is ../../node_modules/@react-native/codegen
    // codegenDir = file("../../node_modules/@react-native/codegen")
    //   The cli.js file which is the React Native CLI entrypoint. Default is ../../node_modules/react-native/cli.js
    // cliFile = file("../../node_modules/react-native/cli.js")

    /* Variants */
    //   The list of variants to that are debuggable. For those we're going to
    //   skip the bundling of the JS bundle and the assets. By default is just 'debug'.
    //   If you add flavors like lite, prod, etc. you'll have to list your debuggableVariants.
    // debuggableVariants = ["liteDebug", "prodDebug"]

    /* Bundling */
    //   A list containing the node command and its flags. Default is just 'node'.
    // nodeExecutableAndArgs = ["node"]
    //
    //   The command to run when bundling. By default is 'bundle'
    // bundleCommand = "ram-bundle"
    //
    //   The path to the CLI configuration file. Default is empty.
    // bundleConfig = file(../rn-cli.config.js)
    //
    //   The name of the generated asset file containing your JS bundle
    // bundleAssetName = "MyApplication.android.bundle"
    //
    //   The entry file for bundle generation. Default is 'index.android.js' or 'index.js'
    // entryFile = file("../js/MyApplication.android.js")
    //
    //   A list of extra flags to pass to the 'bundle' commands.
    //   See https://github.com/react-native-community/cli/blob/main/docs/commands.md#bundle
    // extraPackagerArgs = []

    /* Hermes Commands */
    //   The hermes compiler command to run. By default it is 'hermesc'
    // hermesCommand = "$rootDir/my-custom-hermesc/bin/hermesc"
    //
    //   The list of flags to pass to the Hermes compiler. By default is "-O", "-output-source-map"
    // hermesFlags = ["-O", "-output-source-map"]

    /* Autolinking */
    autolinkLibrariesWithApp()
}

/**
 * Set this to true to Run Proguard on Release builds to minify the Java bytecode.
 */
def enableProguardInReleaseBuilds = false


def jscFlavor = 'org.webkit:android-jsc:+'

android {
    ndkVersion rootProject.ext.ndkVersion
    buildToolsVersion rootProject.ext.buildToolsVersion
    compileSdk rootProject.ext.compileSdkVersion

    namespace "com.app"
    defaultConfig {
        applicationId "com.app"
        minSdkVersion rootProject.ext.minSdkVersion
        targetSdkVersion rootProject.ext.targetSdkVersion
        versionCode 1
        versionName "1.0"
    }
    signingConfigs {
        debug {
            storeFile file('debug.keystore')
            storePassword 'android'
            keyAlias 'androiddebugkey'
            keyPassword 'android'
        }
    }
    buildTypes {
        debug {
            signingConfig signingConfigs.debug
        }
        release {
            // Caution! In production, you need to generate your own keystore file.
            // see https://reactnative.dev/docs/signed-apk-android.
            signingConfig signingConfigs.debug
            minifyEnabled enableProguardInReleaseBuilds
            proguardFiles getDefaultProguardFile("proguard-android.txt"), "proguard-rules.pro"
        }
    }
}

project.ext.vectoricons = [
        iconFontNames: [ 'FontAwesome5_Regular.ttf','Ionicons.ttf','Fontisto.ttf',"AntDesign.ttf",'Entypo.ttf',"FontAwesome.ttf","Foundation.ttf",'Octicons.ttf',
             'FontAwesome5_Solid.ttf', 'FontAwesome5_Brands.ttf','EvilIcons.ttf','Feather.ttf','MaterialIcons.ttf','MaterialCommunityIcons.ttf',
                'SimpleLineIcons.ttf','Zocial.ttf','FontAwesome6_Solid.ttf', 'FontAwesome6_Brands.ttf','FontAwesome6_Regular.ttf'] // <-- add this line (1/2)
]

apply from: file("../../node_modules/react-native-vector-icons/fonts.gradle")

dependencies {
    // The version of react-native is set by the React Native Gradle Plugin
    implementation("com.facebook.react:react-android")

    if (hermesEnabled.toBoolean()) {
        implementation("com.facebook.react:hermes-android")
    } else {
        implementation jscFlavor
    }
}
`
var store = `import { configureStore } from '@reduxjs/toolkit';
import { useDispatch as useAppDispatch, useSelector as useAppSelector } from 'react-redux';
import { persistStore, persistReducer } from 'redux-persist';
import { rootPersistConfig, rootReducer } from './root_reducers.js';

const store = configureStore({
    reducer: persistReducer(rootPersistConfig, rootReducer),
    middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
        serializableCheck: false,
        immutableCheck: false,
    })
});

const persistor = persistStore(store);
const { dispatch } = store;
const useSelector = useAppSelector;
const useDispatch = () => useAppDispatch();
export {store, persistor, dispatch,useSelector,useDispatch};`

var rootRed = `import { combineReducers } from 'redux';
import AsyncStorage from '@react-native-async-storage/async-storage';
import tempReducer from './slices/temp';

const rootPersistConfig = {
  key: 'root',
  storage : AsyncStorage,
  keyPrefix: 'redux-',
  //   whitelist: [],
  //   blacklist: [],
};

const rootReducer = combineReducers({
  temp : tempReducer
});

export { rootPersistConfig,rootReducer};`

var tempSlice = `import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    temp : {}
}

const slice = createSlice({
    name : "temp",
    initialState,
    reducers : {
        setTemp : (state,action) => {
            state.temp = action.payload
        },
        
        
    }
}) 

export const {setTemp} = slice.actions;
export default slice.reducer;`

var tsConfig = `{
  "extends": "@react-native/typescript-config/tsconfig.json",
  "compilerOptions": {
    "target": "es6",
    "module": "esnext",
    "jsx": "react-native",
    "allowSyntheticDefaultImports": true,
    "esModuleInterop": true,  // Add this line
    "skipLibCheck": true,
    "noEmit": true
  } 
}
`

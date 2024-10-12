package cmd

func backendCommands(projName string) []commandType {
	if projName == "" {
		projName = "server"
	}
	return []commandType{
		{command: projName, typ: "mkdir"},
		{command: "/" + projName, typ: "cd"},
		{command: "npm init -y", typ: "exec"},
		{command: "npm install express", typ: "exec"},
		{command: "npm install dotenv", typ: "exec"},
		{command: "npm install cors", typ: "exec"},
		{command: "npm install jsonwebtoken", typ: "exec"},
		{command: "npm install mongoose", typ: "exec"},
		{command: "models", typ: "mkdir"},
		{command: "controllers", typ: "mkdir"},
		{command: "routes", typ: "mkdir"},
		{command: "config", typ: "mkdir"},
		{command: ".env", typ: "write", content: serverEnv},
		{command: "/" + projName + "/config", typ: "cd"},
		{command: "database.js", typ: "write", content: database},
		{command: "/" + projName, typ: "cd"},
		{command: ".gitignore", typ: "write", content: gitIgn},
		{command: "index.js", typ: "write", content: serverData},
	}
}

func reactNativeCommands(projName string) []commandType {
	if projName == "" {
		projName = "MyProject"
	}
	return []commandType{
		{command: "npm i -g @react-native-community/cli", typ: "exec"},
		{command: "npx @react-native-community/cli init " + projName, typ: "exec"},
		{command: "/" + projName, typ: "cd"},
		{command: "npm install @types/react-native", typ: "exec"},
		{command: "tsconfig.json", typ: "write", content: tsConfig},
		{command: "eslint.config.js", typ: "write", content: esConfig},
		// nativewind
		{command: "npm install nativewind", typ: "exec"},
		{command: "npm install --save-dev tailwindcss@3.3.2", typ: "exec"},
		{command: "npx tailwind init", typ: "exec"},
		{command: "tailwind.config.js", typ: "write", content: tWConfig},
		{command: ".eslintrc.js", typ: "write", content: eslint},
		{command: "babel.config.js", typ: "write", content: tWBabel},
		{command: "nativewind-env.d.ts", typ: "write", content: nenvdts},
		// navigation
		{command: "npm install @react-navigation/bottom-tabs", typ: "exec"},
		{command: "npm install react-native-screens", typ: "exec"},
		{command: "npm install @react-navigation/native-stack", typ: "exec"},
		{command: "npm install @react-navigation/native", typ: "exec"},
		{command: "npm install @react-navigation/drawer", typ: "exec"},
		{command: "npm install react-native-safe-area-context", typ: "exec"},
		// packages
		{command: "npm install @react-native-async-storage/async-storage", typ: "exec"},
		{command: "npm install dotenv", typ: "exec"},
		{command: "npm install react-native-dotenv", typ: "exec"},
		{command: "npm install react-native-gesture-handler", typ: "exec"},
		{command: "npm install react-native-reanimated", typ: "exec"},
		{command: "babel.config.js", typ: "write", content: babelFull},
		{command: "env.d.ts", typ: "write", content: envdts},
		{command: "npm install react-native-safe-area-context", typ: "exec"},
		{command: "npm install react-native-vector-icons", typ: "exec"},
		{command: "/" + projName + "/android/app", typ: "cd"},
		{command: "build.gradle", typ: "write", content: buildGradel},
		{command: "/" + projName, typ: "cd"},
		// reduc
		{command: "npm install react-redux", typ: "exec"},
		{command: "npm install redux-persist", typ: "exec"},
		{command: "npm install @reduxjs/toolkit", typ: "exec"},
		{command: "src", typ: "mkdir"},
		{command: "/" + projName + "/src", typ: "cd"},
		{command: "index.tsx", typ: "write", content: ""},
		{command: "redux", typ: "mkdir"},
		{command: "/" + projName + "/src/redux", typ: "cd"},
		{command: "store.js", typ: "write", content: store},
		{command: "root_reducers.js", typ: "write", content: rootRed},
		{command: "slices", typ: "mkdir"},
		{command: "/" + projName + "/src/redux/slices", typ: "cd"},
		{command: "temp.js", typ: "write", content: tempSlice},

		{command: "/" + projName, typ: "cd"},
	}
}

func reactCommands(projName string) []commandType {
	if projName == "" {
		projName = "myproject"
	}
	return []commandType{
		// base
		{command: "npm i -g create-vite", typ: "exec"},
		{command: "npx create-vite " + projName + " --template react", typ: "exec"},
		{command: "/" + projName, typ: "cd"},
		{command: "npm install", typ: "exec"},
		// tailwind
		{command: "npm install -D tailwindcss postcss autoprefixer", typ: "exec"},
		{command: "npx tailwindcss init -p", typ: "exec"},
		{command: "tailwind.config.js", typ: "write", content: tWConfigReact},
		{command: "/" + projName + "/src", typ: "cd"},
		{command: "index.css", typ: "prepend", content: twHeaders},
		{command: "App.css", typ: "prepend", content: twHeaders},
		{command: "/" + projName, typ: "cd"},
		{command: "vercel.json", typ: "write", content: vercelJson},

		// packages
		{command: "npm install react-router-dom", typ: "exec"},
		{command: "npm install axios", typ: "exec"},
		{command: "npm install framer-motion", typ: "exec"},
		{command: "npm install react-icons", typ: "exec"},
		{command: "/" + projName + "/src", typ: "cd"},
		{command: "motionUtils", typ: "mkdir"},
		{command: "/" + projName + "/src/motionUtils", typ: "cd"},
		{command: "motion.js", typ: "write", content: motionReact},
		{command: "/" + projName, typ: "cd"},
		// redux
		{command: "npm install react-redux", typ: "exec"},
		{command: "npm install redux-persist", typ: "exec"},
		{command: "npm install @reduxjs/toolkit", typ: "exec"},
		{command: "/" + projName + "/src", typ: "cd"},
		{command: "redux", typ: "mkdir"},
		{command: "/" + projName + "/src/redux", typ: "cd"},
		{command: "store.js", typ: "write", content: store},
		{command: "root_reducers.js", typ: "write", content: rootRedReact},
		{command: "slices", typ: "mkdir"},
		{command: "/" + projName + "/src/redux/slices", typ: "cd"},
		{command: "temp.js", typ: "write", content: tempSlice},
		{command: "/" + projName, typ: "cd"},

		// UI packages

		{command: "/" + projName, typ: "cd"},
	}
}

func mernCommands(projName string) ([]commandType, []commandType) {
	if projName == "" {
		projName = "myproject"
	}
	reactComm := []commandType{
		// base
		{command: "npm i -g create-vite", typ: "exec"},
		{command: "npx create-vite client --template react", typ: "exec"},
		{command: "/" + projName + "/client", typ: "cd"},
		{command: "npm install", typ: "exec"},
		// tailwind
		{command: "npm install -D tailwindcss postcss autoprefixer", typ: "exec"},
		{command: "npx tailwindcss init -p", typ: "exec"},
		{command: "tailwind.config.js", typ: "write", content: tWConfigReact},
		{command: "/" + projName + "/client/src", typ: "cd"},
		{command: "index.css", typ: "prepend", content: twHeaders},
		{command: "App.css", typ: "prepend", content: twHeaders},
		{command: "/" + projName + "/client", typ: "cd"},
		{command: "vercel.json", typ: "write", content: vercelJson},

		// packages
		{command: "npm install react-router-dom", typ: "exec"},
		{command: "npm install axios", typ: "exec"},
		{command: "npm install framer-motion", typ: "exec"},
		{command: "npm install react-icons", typ: "exec"},
		{command: "/" + projName + "/client/src", typ: "cd"},
		{command: "motionUtils", typ: "mkdir"},
		{command: "/" + projName + "/client/src/motionUtils", typ: "cd"},
		{command: "motion.js", typ: "write", content: motionReact},
		{command: "/" + projName + "/client", typ: "cd"},
		// redux
		{command: "npm install react-redux", typ: "exec"},
		{command: "npm install redux-persist", typ: "exec"},
		{command: "npm install @reduxjs/toolkit", typ: "exec"},
		{command: "/" + projName + "/client/src", typ: "cd"},
		{command: "redux", typ: "mkdir"},
		{command: "/" + projName + "/client/src/redux", typ: "cd"},
		{command: "store.js", typ: "write", content: store},
		{command: "root_reducers.js", typ: "write", content: rootRedReact},
		{command: "slices", typ: "mkdir"},
		{command: "/" + projName + "/client/src/redux/slices", typ: "cd"},
		{command: "temp.js", typ: "write", content: tempSlice},
		{command: "/" + projName + "/client", typ: "cd"},

		// UI packages

		{command: "/" + projName + "/client", typ: "cd"},
	}
	backComm := []commandType{
		{command: "server", typ: "mkdir"},
		{command: "/" + projName + "/server", typ: "cd"},
		{command: "npm init -y", typ: "exec"},
		{command: "npm install express", typ: "exec"},
		{command: "npm install dotenv", typ: "exec"},
		{command: "npm install cors", typ: "exec"},
		{command: "npm install jsonwebtoken", typ: "exec"},
		{command: "npm install mongoose", typ: "exec"},
		{command: "models", typ: "mkdir"},
		{command: "controllers", typ: "mkdir"},
		{command: "routes", typ: "mkdir"},
		{command: "config", typ: "mkdir"},
		{command: ".env", typ: "write", content: serverEnv},
		{command: "/" + projName + "/server/config", typ: "cd"},
		{command: "database.js", typ: "write", content: database},
		{command: "/" + projName + "/server", typ: "cd"},
		{command: ".gitignore", typ: "write", content: gitIgn},
		{command: "index.js", typ: "write", content: serverData},
	}
	return reactComm, backComm
}

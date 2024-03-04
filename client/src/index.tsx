/* @refresh reload */
import { render } from "solid-js/web";

import "./index.css";
import { Route, Router } from "@solidjs/router";
import { ParentProps, lazy } from "solid-js";
import Header from "./components/Header";
import { AuthContextProvider } from "./context/auth.context";
import { PostsContextProvider } from "./context/posts.context";

const Home = lazy(() => import("./routes/home"));
const Login = lazy(() => import("./routes/login"));
const Register = lazy(() => import("./routes/register"));
const CreatePost = lazy(() => import("./routes/create-post"));
const PostPage = lazy(() => import("./routes/post-page"));

const root = document.getElementById("root");

const App = (props: ParentProps) => (
	<div>
		<Header />
		{props.children}
	</div>
);

render(
	() => (
		<AuthContextProvider>
			<PostsContextProvider>
				<Router root={App}>
					<Route path={"/"} component={Home} />
					<Route path={"/login"} component={Login} />
					<Route path={"/register"} component={Register} />
					<Route path={"/create-post"} component={CreatePost} />
					<Route path={"/post/:id"} component={PostPage} />
				</Router>
			</PostsContextProvider>
		</AuthContextProvider>
	),
	root!
);

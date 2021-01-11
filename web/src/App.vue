<!-- @format -->

<template>
	<div id="app">
		<div id="nav">
			<router-link to="/"
				><b-icon pack="fas" icon="home" size="is-small"> </b-icon>
				Home</router-link
			>
		</div>
		<router-view @refresh-state="refreshState" />
	</div>
</template>

<style>
#app {
	text-align: center;
	color: #2c3e50;
}
#nav {
	padding: 30px;
}
#nav a {
	font-weight: bold;
	color: #2c3e50;
}
#nav a.router-link-exact-active {
	color: #42b983;
}
</style>

<script>
export default {
	data() {
		return {
			logged: localStorage.getItem("x_api_key") != null,
		};
	},
	methods: {
		refreshState() {
			this.logged = localStorage.getItem("x_api_key") != null;
		},
	},
	mounted() {},
};
</script>

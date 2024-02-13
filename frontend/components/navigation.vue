<template>
    <v-app-bar id="navigation" color="primary" :height="75" :elevation="0">
        <div id="nav_left">
            <h1 id="nav_brandname">MusicStreaming</h1>
            <nav class="ms-3">
                <v-btn href="/" flat>
                    <v-icon start icon="mdi-home"></v-icon>
                    Home
                </v-btn>
                <v-btn flat>
                    <v-icon start icon="mdi-library"></v-icon>
                    Library
                </v-btn>
            </nav>
        </div>

        <div id="nav_right">
            <v-btn class="me-5" flat>
                <v-icon start icon="mdi-cloud-upload"></v-icon>
                Upload
            </v-btn>

            <v-avatar color="secondary" size="large">
                <v-img alt="Avatar" src="https://avatars.githubusercontent.com/u/72092675?v=4"></v-img>
            </v-avatar>

            <v-btn class="ms-3" variant="tonal" icon="mdi-dots-horizontal"></v-btn>
        </div>
    </v-app-bar>
</template>

<style lang="scss">
#navigation {
    padding: 0 80px;

    #nav_brandname {
        position: relative;
        top: 2px;
        font-family: "sans-serif";
        font-size: 20px;
    }

    #nav_left {
        display: flex;
    }

    #nav_right {
        display: inline-block;
        margin-left: auto;
    }
}
</style>

<script>
export default {
    name: "Navigation"
}
</script>

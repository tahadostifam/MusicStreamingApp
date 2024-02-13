<template>
    <v-container>
        <v-card id="signin_form" elevation="0">
            <v-toolbar color="primary" :dark="true">
                <v-toolbar-title>Signin</v-toolbar-title>
            </v-toolbar>
            <v-form class="py-3 ">
                <v-text-field flat label="Email" variant="solo" placeholder="Enter your email..."></v-text-field>
                <v-checkbox color="primary" label="Remember me"></v-checkbox>
                <v-btn size="large" color="primary" variant="tonal" block>Signin</v-btn>
            </v-form>
        </v-card>
    </v-container>
</template>

<style>
#signin_form {
    position: relative;
    left: 50%;
    transform: translate(-50%);
    margin-top: 100px;
    width: 550px;

    @media screen and (max-width: 1000px) {
        & {
            width: 90%;
            margin-top: 20px;
        }
    }
}
</style>

<script>
</script>

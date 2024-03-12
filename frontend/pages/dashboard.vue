<template>
    <div class="dashboard_contianer">
        <v-sheet class="py-10 px-8" :width="'100%'" color="grey-lighten-3">
            <div class="d-flex align-center">
                <v-avatar color="secondary" size="200">
                    <v-img alt="Avatar" src="https://avatars.githubusercontent.com/u/72092675?v=4"></v-img>
                </v-avatar>
                <h1 class="ml-6 mb-15 bg-grey px-5">Taha</h1>
            </div>
        </v-sheet>

        <div class="my-10">
            <h1>Your Tracks</h1>

            <v-container>
                <v-row>
                    <v-col class="pa-0" cols="8">
                        <div v-for="n in 15" :key="n" class="d-flex align-center">
                            <v-avatar class="rounded" size="65">
                                <v-img src="https://cdn.vuetifyjs.com/images/john.jpg" alt="John"></v-img>
                            </v-avatar>
                            <v-col class="ml-1">
                                <h3>The Weeknd</h3>
                                <div>
                                    <span class="text-caption text-grey">
                                        <v-icon icon="mdi-play"></v-icon>
                                        2000
                                    </span>
                                    <span class="text-caption text-grey ml-2">
                                        <v-icon icon="mdi-heart"></v-icon>
                                        300+
                                    </span>
                                    <span class="text-caption text-grey ml-2">
                                        <v-icon icon="mdi-comment"></v-icon>
                                        25+
                                    </span>
                                </div>
                            </v-col>
                        </div>
                    </v-col>

                    <v-col cols="4" class="pt-15">
                        <info-sidebar></info-sidebar>
                    </v-col>
                </v-row>
            </v-container>
        </div>
    </div>
</template>

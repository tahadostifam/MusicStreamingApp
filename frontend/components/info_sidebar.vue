<template>
    <div>
        <div>
            <div class="d-flex justify-space-between">
                <h4 class="font-weight-thin mb-3">Famous Artists</h4>
                <a href="/">View All</a>
            </div>

            <div>
                <v-slide-group v-model="recentlyPlayed">
                    <v-slide-group-item v-for="n in 15" :key="n" v-slot="{ isSelected, toggle, selectedClass }">
                        <v-card :ripple="false" flat :class="['ma-1', 'd-flex flex-column align-center', selectedClass]"
                            @click="toggle">
                            <v-avatar size="65">
                                <v-img src="https://cdn.vuetifyjs.com/images/john.jpg" alt="John"></v-img>
                            </v-avatar>
                            <span class="mt-2 text-caption text-truncate">The Weeknd</span>
                        </v-card>
                    </v-slide-group-item>
                </v-slide-group>
            </div>
        </div>

        <div class="mt-10">
            <div class="d-flex justify-space-between">
                <h4 class="font-weight-thin mb-3">Artists you should follow</h4>
                <a href="/">Refresh all</a>
            </div>

            <div>
                <v-card flat>
                    <v-container class="align-center justify-space-between pa-0">
                        <div v-for="n in 5" :key="n" class="d-flex align-center">
                            <v-avatar size="65">
                                <v-img src="https://cdn.vuetifyjs.com/images/john.jpg" alt="John"></v-img>
                            </v-avatar>
                            <v-col class="ml-1">
                                <h3>The Weeknd</h3>
                                <div class="d-flex justify-space-between align-center">
                                    <div>
                                        <span class="text-caption text-grey">
                                            <v-icon icon="mdi-account-circle"></v-icon>
                                            200k
                                        </span>
                                        <span class="text-caption text-grey ml-2">
                                            <v-icon icon="mdi-music-circle"></v-icon>
                                            20
                                        </span>
                                    </div>
                                    <div class="text-right">
                                        <v-btn color="grey" size="small" class="text-none"
                                            prepend-icon="mdi-account-plus" variant="text">
                                            Follow
                                        </v-btn>
                                    </div>
                                </div>
                            </v-col>
                        </div>
                    </v-container>
                </v-card>
            </div>
        </div>

        <div class="mt-10">
            <div class="d-flex justify-space-between">
                <h4 class="font-weight-thin mb-3">1253 Likes</h4>
                <a href="/">View All</a>
            </div>

            <div>
                <v-card flat>
                    <v-container class="align-center justify-space-between pa-0">
                        <div v-for="n in 5" :key="n" class="d-flex align-center">
                            <v-avatar class="rounded-0" size="65">
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
                    </v-container>
                </v-card>
            </div>
        </div>
    </div>
</template>

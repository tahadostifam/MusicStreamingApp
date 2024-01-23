<template>
    <v-container>
        <v-row>
            <v-col cols="9">
                <div class="mb-5">
                    <h2 class="font-weight-black mb-3">Recently Played</h2>
                    <v-slide-group v-model="recentlyPlayed">
                        <v-slide-group-item v-for="n in 15" :key="n" v-slot="{ isSelected, toggle, selectedClass }">
                            <v-card color="grey-lighten-1" :class="['ma-1', selectedClass]" height="200" width="200"
                                @click="toggle">
                                <div class="d-flex fill-height align-center justify-center">
                                    <v-img cover width="100%" height="100%"
                                        src="https://picsum.photos/500/300?image=232"></v-img>
                                </div>
                            </v-card>
                        </v-slide-group-item>
                    </v-slide-group>
                </div>

                <div class="mb-5">
                    <h2 class="font-weight-black mb-3">More of what you like</h2>
                    <v-slide-group v-model="recentlyPlayed">
                        <v-slide-group-item v-for="n in 15" :key="n" v-slot="{ isSelected, toggle, selectedClass }">
                            <v-card color="grey-lighten-1" :class="['ma-1', selectedClass]" height="200" width="200"
                                @click="toggle">
                                <div class="d-flex fill-height align-center justify-center">
                                    <v-img cover width="100%" height="100%"
                                        src="https://picsum.photos/500/300?image=232"></v-img>
                                </div>
                            </v-card>
                        </v-slide-group-item>
                    </v-slide-group>
                </div>
            </v-col>

            <v-col cols="3" class="pt-15">
                <div>
                    <div class="d-flex justify-space-between">
                        <h4 class="font-weight-thin mb-3">Famous Artists</h4>
                        <a href="/">View All</a>
                    </div>

                    <div>
                        <v-slide-group v-model="recentlyPlayed">
                            <v-slide-group-item v-for="n in 15" :key="n" v-slot="{ isSelected, toggle, selectedClass }">
                                <v-card :ripple="false" flat
                                    :class="['ma-1', 'd-flex flex-column align-center', selectedClass]" @click="toggle">
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
            </v-col>
        </v-row>
    </v-container>
</template>

<style>
.v-slide-group {
    position: relative;

    .v-slide-group__prev,
    .v-slide-group__next {
        position: absolute;
        top: 50%;
        transform: translate(0, -50%);
        z-index: 2;
        padding: 50px 0;
        color: #fff;
        font-size: 28px;
        opacity: 0.97;
    }

    .v-slide-group__prev {
        left: 0;
    }

    .v-slide-group__next {
        right: 0;
    }

    .v-slide-group__prev.v-slide-group__prev--disabled {
        opacity: 0 !important;
    }

    .v-slide-group__next.v-slide-group__next--disabled {
        opacity: 0 !important;
    }
}
</style>

<script>
export default {
    name: "index",
    data() {
        return {
            recentlyPlayed: null
        }
    },
}
</script>

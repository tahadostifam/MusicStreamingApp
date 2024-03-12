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
                <info-sidebar></info-sidebar>
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

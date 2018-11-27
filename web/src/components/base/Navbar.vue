<template>
    <v-navigation-drawer id="nav-bar"
      persistent
      enable-resize-watcher
      fixed
      app
      permanent
      :mini-variant="value"
      :clipped="clipped"
    >
    <slot></slot>
    </v-navigation-drawer>
</template>

<script>
export default {
    props: {
        value: {
        type: Boolean
        },
        mini: {
        type: Boolean
        },
        clipped: {
        type: Boolean
        }
    }
}
</script>

<style scoped>
#nav-bar {
  z-index: 2;
}
</style>
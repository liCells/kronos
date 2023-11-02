<script setup>
const props = defineProps({
  selectedPlugin: String
})

</script>

<template>
  <div>
    <h1>{{ props.selectedPlugin }}</h1>
  </div>
</template>

<style scoped>
</style>

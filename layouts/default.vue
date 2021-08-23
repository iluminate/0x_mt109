<template>
  <div>
    <Nuxt />
  </div>
</template>

<script>
export default {
  head() {
    return {
      bodyAttrs: {
        class: 'bg-sky'
      }
    }
  }
}
</script>

<style>
.bg-sky {
  background-color: #e2f2ff;
}
</style>


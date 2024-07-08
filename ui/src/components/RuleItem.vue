<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
  <label>
    Type
    <select>
      <option>Request to HTTP Service</option>
    </select>
  </label>
  <label>Result <input type="number" min="100" max="599" /></label>
  <label>
    Request Host
    <select>
      <option>en.wikipedia.org</option>
      <option>google.com</option>
      <option>apple.com</option>
    </select>
  </label>
</template>

<style scoped></style>

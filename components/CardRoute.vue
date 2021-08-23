<template>
  <div class="maptag-row">
    <span>{{ title }}</span>
    <br/>
    <p :class="classObject">{{ data || '-' }}</p>
  </div>
</template>

<script>
export default {
  name: "CardRoute",
  props: ['title', 'data', 'classObject'],
}
</script>

<style scoped>
.maptag-row {
  display: inline-block;
  padding: 8px 22px;
}

.maptag-row span {
  color: #999999;
  font-size: 0.85em;
}
p {
  margin: 0px;
}
p.blue {
  padding: 4px 12px;
  color: #e2f2ff;
  background-color: #3b8bff;
  -webkit-border-radius: 20px;
  -moz-border-radius: 20px;
  border-radius: 20px;
  font-size: 0.85em;
}

p.green {
  padding: 4px 12px;
  color: #e2f2ff;
  background-color: #52ba5d;
  -webkit-border-radius: 20px;
  -moz-border-radius: 20px;
  border-radius: 20px;
  font-size: 0.85em;
}
</style>
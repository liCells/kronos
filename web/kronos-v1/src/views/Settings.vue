<script setup>
</script>

<template>
  <div>
    <h1>Settings</h1>
    <div class="input">
      <label for="input" class="text">Server Url:</label>
      <input type="text" placeholder="http://localhost:18000" name="input" class="input">
    </div>
  </div>
</template>

<style scoped>
h1 {
  text-align: center;
}

.input {
  position: static;
}

.input label.text {
  display: inline-block;
  width: 7rem;
  text-align: right;
  margin-right: 1rem;
}

.input input[type=text].input {
  height: 2rem;
  width: 15rem;
  font-size: 1rem;
  border: 2px #9a9a9a solid;
  border-radius: 5px;
}

.input input[type=text].input:focus {
  outline: none;
}
</style>

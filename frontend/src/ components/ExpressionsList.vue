<template>
  <div>
    <h2>Expressions List</h2>
    <ul>
      <li v-for="expression in expressions" :key="expression.id">
        {{ expression.text }} - {{ expression.status }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      expressions: [] // Array to hold expressions fetched from backend
    };
  },
  created() {
    // Implement logic to fetch expressions from backend
  }
};
</script>

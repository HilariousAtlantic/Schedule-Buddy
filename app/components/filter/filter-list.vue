<template>

  <div class="filter-list">

    <filter-group filterType="start" label="Start Time">

      <start-filter v-for="filter in getFilters('start')"
        :id="filter.id"
        :options="filter.options"
      ></start-filter>

    </filter-group>

    <filter-group filterType="finish" label="Finish Time">

      <finish-filter v-for="filter in getFilters('finish')"
        :id="filter.id"
        :options="filter.options"
      ></finish-filter>

    </filter-group>

    <filter-group filterType="break" label="Break Time">

      <break-filter v-for="filter in getFilters('break')"
        :id="filter.id"
        :options="filter.options"
      ></break-filter>

    </filter-group>

    <filter-group filterType="class" label="Class Load">

      <class-filter v-for="filter in getFilters('class')"
        :id="filter.id"
        :options="filter.options"
      ></class-filter>

    </filter-group>

    <button class="filter-button btn btn-danger"
      :class="{disabled: !canClearFilters}"
      @click="clearFilters">

      <span>Remove Filters</span>

    </button>

    <button class="filter-button btn btn-primary"
      :class="{disabled: !canApplyFilters}"
      @click="applyFilters">

      <span>Apply Changes</span>

    </button>

  </div>

</template>

<script>

  import StartFilter from './start-filter.vue';
  import FinishFilter from './finish-filter.vue';
  import BreakFilter from './break-filter.vue';
  import ClassFilter from './class-filter.vue';
  import SectionFilter from './section-filter.vue';
  import FilterGroup from './filter-group.vue';

  export default {

    name: 'filter-list',

    components: {StartFilter, FinishFilter, BreakFilter, ClassFilter, SectionFilter, FilterGroup},

    data() {

      return {

        editing: ''

      }

    },

    computed: {

      canClearFilters() {

        return this.$store.state.filters.length > 0;

      },

      canApplyFilters() {

        for (let filter of this.$store.state.filters) {

          if (filter.changes) return true;

        }

        return false;

      }

    },

    methods: {

      selectEditor(type) {

        this.editing = this.editing == type ? '' : type;

      },

      getFilters(type) {

        return this.$store.state.filters.filter(filter => filter.type === type);

      },

      clearFilters() {

        if (this.canClearFilters) {

          this.$store.dispatch('clearFilters');

        }

      },

      applyFilters() {

        if (this.canApplyFilters) {

          this.$store.dispatch('applyFilters');

        }

      }

    }

  }

</script>

<style scoped>

  .filter-list {
    display: flex;
  }

  .filter-group {
    flex: 1;
  }

  .filter-group + .filter-group, .filter-button {
    margin-left: 5px;
  }

  .filter {
    margin-bottom: 10px;
  }

</style>

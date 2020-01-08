<template>
  <div class="course-box" @click="loadModal">
    <div class="course-info">
      <div class="subject-sort" v-html="course.kamoku_sort"></div>
      <div class="title-memo-en" v-html="course.title_memo"></div>
      <!-- <div class="language-en" v-html="course.Language_en"></div> -->
    </div>
    <div class="title-en" v-html="course.title"></div>
    <div class="semester-en" v-html="course.semester"></div>
    <div class="days-en" v-html="course.class_day"></div>
    <!-- <div class="faculty-in-charge" v-html="course.Faculty_in_charge"></div> -->
  </div>
</template>

<style>
.course-box {
  border-radius: 25px;
  padding: 20px;
  width: 25em;
  margin: 1%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  line-height: 2em;
  box-shadow: 0px 2px 20px 0px #e5ddf5;
}

.course-box:hover {
  cursor: pointer;
  background: Highlight;
}

.course-box .highlight {
  color: red;
}

.course-info {
  display: inline-flex;
  justify-content: space-evenly;
}

.subject-sort {
  float: left;
}

.language-en {
  float: right;
}
</style>
<style class="dark">
body.dark .course-box {
  color: white;
  border-radius: 25px;
  padding: 20px;
  width: 25em;
  margin: 1%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  line-height: 2em;
  box-shadow: 0px 2px 20px 0px #e5ddf5;
}
body.dark .course-box:hover {
  cursor: pointer;
  color: white;
  background-color: rgb(20, 31, 39);
}
body.dark .course-box .highlight {
  color: #ff6200;
}
</style>
<script>
export default {
  data() {
    return {
      showModal: false
    };
  },
  props: ["searchResult", "query"],
  computed: {
    course() {
      return this.searchResult;
    }
  },
  methods: {
    loadModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    }
  }
};
</script>

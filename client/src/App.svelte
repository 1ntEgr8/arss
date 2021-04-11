<style>
    main {
        display: grid;
        grid-template-columns: 0.2fr 1fr;
        grid-gap: var(--spacing-big);
        margin: 0px var(--spacing-big);
    }
</style>

<script>
  import Header from "./Header.svelte";
  import Sidebar from "./Sidebar.svelte";
  import Feed from "./Feed.svelte";
</script>

<Header />
<main>
  <Sidebar/>
  <Feed source="all"/>
</main>

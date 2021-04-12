<style>
    .feed {
        margin: 0px var(--spacing);
    }
</style>

<script>
    import AppBar from "./AppBar.svelte";
    import FeedItem from "./FeedItem.svelte";
    import Button from "./Button.svelte";
    
    export let source = "[no source]";

    async function fetchFeed() {
      let res = await fetch("http://localhost:8000/feed");
      let data = await res.json();
      return data["items"];
    }

    const promise = fetchFeed();
</script>

<div class="feed">
    <AppBar>
        <div slot="main">
          <span class="font-size-l">feed</span>
          <span>/{source}</span>
        </div>
        <div class="font-size-s" slot="options">
          <Button text="sort by" />
        </div>
    </AppBar>
    {#await promise}
      loading...
    {:then feedItems}
        {#each feedItems as item}
            <FeedItem {...item} />
        {/each}
    {/await}
</div>


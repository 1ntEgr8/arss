<style>
    .feed {
        margin: 0px var(--spacing);
    }
</style>

<script>
    import AppBar from "./AppBar.svelte";
    import FeedItem from "./FeedItem.svelte";
    import Button from "./Button.svelte";
    
    export let fetchFeed;
    export let source = "[no source]";
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
    {#await fetchFeed}
      loading...
    {:then feedItems}
        {#each feedItems as { 
          title = "[no title]", 
          link = "/", 
          published = "[no date]", 
          categories = [] 
        }}
            <FeedItem {title} {link} {published} {categories} />
        {/each}
    {/await}
</div>


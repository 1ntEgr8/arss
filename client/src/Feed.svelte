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

    const feedItems = [
        {
            title: "hi",
            link: "https://www.example.com",
            source: "lobsters",
            categories: ["test-1", "test-2", "test-3"],
            comments: ""
        },
        {
            title: "there",
            link: "https://www.example.com",
            source: "lobsters",
            categories: ["test-1", "test-2", "test-3"],
            comments: ""
        }
    ];
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
    {#each feedItems as item}
        <FeedItem {...item} />
    {/each}
</div>


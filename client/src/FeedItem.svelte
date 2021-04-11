<style>
    .feed-item {
        padding: var(--spacing) 0px;
    }

    .title {
        display: block;
    }

    .category {
        background: #fffcd7;
        border: 1px solid #d5d458;
        border-radius: 2px;
        padding: 0px 0.4em 1px 0.4em;
        margin: 0px 2px;
    }
</style>

<script>
    export let title = "[no title]";
    export let link = "https://www.example.com";
    export let source = "[no source]";
    export let categories;
    export let comments;
</script>

<div class="feed-item">
    <a class="title font-size-l" href={link} target="_blank"> 
        {title}
    </a>
    <span class="font-size-s">
        <span class="bold">from: </span>
        {source}
    </span>
    {#each categories as cat}
        <span class="font-size-xs category">{cat}</span> 
    {/each}
    <a class="font-size-s" href={comments} target="_blank">comments</a>
</div>

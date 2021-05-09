<style>
    .feed-item {
        padding: var(--spacing) 0px;
    }

    .title {
        display: block;
    }

    .category {
        background: var(--accent-color-light);
        color: var(--text-color-dark);
        border: 1px solid var(--accent-color-dark);
        border-radius: 2px;
        padding: 0px 0.4em 1px 0.4em;
        margin: 0px 2px;
    }
</style>

<script>
    export let title = "[no title]";
    export let link = "https://www.example.com";
    export let categories = [];
    export let published = "[no date]"
</script>

<div class="feed-item">
    <a class="title font-size-l" href={link} target="_blank"> 
        {title}
    </a>
    {#each categories as cat}
        <span class="font-size-xs category">{cat}</span> 
    {/each}
    <span class="font-size-xs">{published}</span>
</div>

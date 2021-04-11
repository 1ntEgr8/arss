<style>
    .app-bar {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
    }
</style>

<div class="app-bar">
    <slot name="main"></slot>
    <slot name="options"></slot>
</div>

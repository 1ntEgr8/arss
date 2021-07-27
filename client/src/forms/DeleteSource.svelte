<style>
    .form {
        width: 100%;
        margin: var(--spacing-small) 0px;
        display: grid;
        grid-template-columns: repeat(2,1fr);
        grid-gap: var(--spacing-small);
    }

    .btn-container {
        display: grid;
        grid-template-columns: repeat(2,1fr);
        grid-column: 1/3;
        grid-gap: var(--spacing-small);
    }

    #confirm {
        grid-column: 1/3;
    }

    #cancel {
        background: var(--danger-color);
    }

    #cancel:hover {
        background: var(--danger-color-hover);
    }

    button {
        font-size: var(--font-size-s);
    }
</style>

<script>
    export let handleYes = undefined;
    export let handleNo = undefined;
</script>

<form
    class="form font-size-m"
    on:submit|preventDefault={handleYes}
>
    <span id="confirm">are you sure?</span>
    <div class="btn-container">
        <button type="submit"> yes </button>
        <button id="cancel" on:click={handleNo} type="button"> no </button>
    </div>
</form>

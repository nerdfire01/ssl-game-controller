<template>
    <div class="btn-group-toggle btn-group">
        <label v-help-text="leftTitle"
               :class="{btn:true, 'btn-active': (selectedValue === leftValue), 'btn-passive': (selectedValue !== leftValue)}"
               @click="callback(leftValue)"
               :disabled="disabled">
            {{leftLabel}}
        </label>
        <label v-help-text="rightTitle"
               :class="{btn:true, 'btn-active': (selectedValue === rightValue), 'btn-passive': (selectedValue !== rightValue)}"
               @click="callback(rightValue)"
               :disabled="disabled">
            {{rightLabel}}
        </label>
    </div>
</template>

<script>
    export default {
        name: "DualSwitch",
        props: {
            leftTitle: String,
            leftLabel: String,
            leftValue: String,
            rightTitle: String,
            rightLabel: String,
            rightValue: String,
            callback: {
                type: Function,
                default: function () {
                }
            },
            selectedValue: {
                type: String,
                default: ''
            },
            disabled: {
                type: Boolean,
                default: false
            },
        }
    }
</script>

<style scoped>

    .btn-active {
        background-color: green;
        border-color: green;
        color: white;
    }

    .btn-passive {
        background-color: red;
        border-color: red;
        color: white;
    }

</style>

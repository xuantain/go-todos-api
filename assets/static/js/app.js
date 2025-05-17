document.addEventListener('DOMContentLoaded', function(e) {

    $('.btn-del__todo').click(function(e) {
        let todoId = $(this).data('id');
        let text;

        if (confirm('Delete Todo ' + todoId + ' ?')) {
            text = 'You pressed OK!';
            let $form = $('#TodoListForm');
            $form.find('[name=id]').val(todoId);
            $form.submit();
        } else {
            text = 'You canceled!';
        }

        console.log(text);
    });
});

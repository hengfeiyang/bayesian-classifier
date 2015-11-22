$(function () {
	// categorize
	var categorizeForm = $('#categorizeForm');
	var categorizeResult = $('.categorizeResult');

	categorizeForm.find('.btn').click(function () {
		var doc = categorizeForm.find('[name=document]');
		if (doc.val() == '') {
			return;
		}
		$.post('/api/categorize', { "doc": doc.val() }, function (ret) {
			if (ret && ret.code) {
				alert(ret.message);
			} else {
				categorizeResult.empty();
				for (var i in ret.data) {
					categorizeResult.append('<li>' + ret.data[i]['category'] + ' / ' + ret.data[i]['score'] + '</li>');
				}
			}
		}, 'json');
	});
	
	// train
	var trainForm = $('#trainForm');
	var trainResult = $('.trainResult');

	trainForm.find('.btn').click(function () {
		var doc = trainForm.find('[name=document]');
		var category = trainForm.find('input[name=category]');
		if (doc.val() == '' || category.val() == "") {
			alert("内容不能为空");
			return;
		}
		$.post('/api/train', { "doc": doc.val(), "category": category.val() }, function (ret) {
			if (ret && ret.code) {
				trainResult.html(ret.message);
			} else {
				trainResult.html(ret.message || "学习成功");
			}
		}, 'json');
	});
	
	// words score
	var scoreForm = $('#scoreForm');
	var scoreResult = $('.scoreResult');

	scoreForm.find('.btn').click(function () {
		var word = scoreForm.find('input[name=word]');
		var category = scoreForm.find('input[name=category]');
		if (word.val() == '') {
			return;
		}
		$.post('/api/score', { "word": word.val(), "category": category.val() }, function (ret) {
			if (ret && ret.code) {
				alert(ret.message);
			} else {
				scoreResult.empty();
				for (var i in ret.data) {
					scoreResult.append('<li>' + ret.data[i]['category'] + ' / ' + ret.data[i]['score'] + '</li>');
				}
			}
		}, 'json');
	});
	
});
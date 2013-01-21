function SchalmeiCtrl($scope, $http) {
	$scope.show = function(name, val) {
		$scope.shown = name;

		if (name == 'show-ranks') {
			$http.get($('#show-rank-list').attr('data-url')).
				success(function(data, status) {
					$scope.ranks = data;
				});
		}
		else if (name == 'get-rank') {
			$http({
				method: 'GET',
				url: val.Url
			}).success(function(data, status) {
				$scope.rank = data;
			});
		}
	};

	$scope.isActive = function(name) {
		return $scope.shown == name ? 'active' : '';
	};

	$scope.createRank = function() {
		if (!$scope.createRankName)
			return;

		$http.post($('#create-rank-form').attr('action'), $scope.createRankName).
			success(function(data, status) {
				$scope.show('get-rank', data);
			});
	};

	$scope.uploadNote = function() {
		var form = $('#file_form');
		$scope.uploadImageUrl($scope.rank.Url, function(data) {
			var url = JSON.parse(data);
			form.attr('action', url);
			form.ajaxSubmit({
				success: function(data) {
					console.log(data);
					form[0].reset();
				}
			});
		});
	};

	$scope.uploadImageUrl = function(url, callback) {
		$.ajax({
			url: url
		}).done(callback);
	};
}

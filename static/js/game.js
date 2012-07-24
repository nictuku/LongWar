/* LongWar, by Yves Junqueira.
 *
 *
 */

window.onload = (function() {

	Crafty.init();

	// This will fetch "REPONAME/hexmap-VERSION.js".
	Crafty.modules("../js/thirdparty", {
		hexmap: 'release'
	}, longWarMap)
});

const spriteSize = 64;
const tiles = ["grass", "forrest", "swamp", "redMountain", "water", "wood", "sand", "nothing"];
const lastTerrain = 5; // Terrains are from "grass" to "wood".

function longWarMap() {

	const width = 800;
	const height = 600;

	// 30x30 is able to cover a 800x600 screen.
	const numColumns = 30;
	const numLines = 30;
	// Each individual item in the sprite has spriteSize*spriteSize pixels.
	Crafty.sprite(spriteSize, "../images/hexagons.png", {
		// Sprites from 0 to lastTerrain are used for random terrain generation.
		grass: [0, 0],
		forrest: [1, 0],
		swamp: [2, 0],
		redMountain: [3, 0],
		nothing: [4, 0],
		water: [5, 0],
		wood: [6, 0],
		sand: [7, 0]
	});

	var hexmap = Crafty.e("Hexmap").hexmap(spriteSize, spriteSize);

	// OH THE HORROR.
	hexmap.loadMap = function(mapWidth, mapHeight, mapFromServer, creationFunc) {
		// mapFromServer is a 30*30 int array.
		this._mapWidth = mapWidth;
		this._mapHeight = mapHeight;
		for (j = 0; j < this._mapHeight; j++) {
			for (i = 0; i < this._mapWidth; i++) {
				t = -1;
				pos = j * (this._mapHeight) + i;
				if ((mapFromServer) && (mapFromServer[pos])) {
					t = mapFromServer[pos]
				}
				if (t == -1) {
					t = 0;
				}
				tile = creationFunc(t);

				tile.addComponent("HexmapNode").hexmapNode(i, j);
				this.placeTile(i, j, tile).setTile(i, j, tile);
			}
		}
	}

	$.ajax({
		url: 'http://localhost:8080/rnd?jsoncallback=?',
		async: true,
		dataType: 'jsonp',
		timeout: 3000,

		error: function(xhr, textStatus) {
			// TODO: Show an error on the web page. The error is usually a
			// timeout, because jQuery can't run this callback properly for
			// jsonp.
			console.log("error:", textStatus);
		},
		success: function(data) {

			// Fetch a random map from the server.
			var mapFromServer = [];
			$.each(data, function(i, v) {
				mapFromServer.push(v);
			});

			hexmap.loadMap(numColumns, numLines, mapFromServer, function(tileIndex) {
				var terrain = tiles[tileIndex];
				return Crafty.e("2D, DOM, Mouse, " + terrain).attr({
					tileIndex: tileIndex
				}).bind("Click", function() {
					var nextTile = (this.tileIndex + 1) % (lastTerrain + 1);
					this.removeComponent(tiles[this.tileIndex])
					this.addComponent(tiles[nextTile])
					this.tileIndex = nextTile;
				}).bind("MouseOver", function() {
				/*
					if (grass) {
						if (grass) grass.forEach(function(tile) {
						tile.removeComponent("grass").addComponent("forrest");
					});

				grass = hexmap.findpath(grass, this);
				if (grass) grass.forEach(function(tile) {
					tile.removeComponent("forrest").addComponent("grass");
				});
					}
				*/
				});
			});


		}
	});
};
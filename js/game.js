/* LongWar, by Yves Junqueira.
 *
 *
 */

window.onload = (function() {

	Crafty.init();

	// It fetches "REPONAME/hexmap-VERSION.js".
	Crafty.modules("../js/thirdparty", {
		hexmap: 'release'
	}, longWarMap)
});

spriteSize = 64;
var tiles = ["grass", "forrest", "swamp", "redMountain", "water", "wood", "sand", "nothing"]
// Later this will be a subset of the above.
lastTerrain = 5; //tiles.length-1;

function longWarMap() {

	width = 800;
	height = 600;

	// 30x30 is able to cover an 800x600 screen.
	numColumns = 30;
	numLines = 30;
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
		// mapFromServer is a two-dimensional array pointing to tiles with the following attributes:
		// - terrain
		this._mapWidth = mapWidth;
		this._mapHeight = mapHeight;
		for (j = 0; j < this._mapHeight; j++) {
			for (i = 0; i < this._mapWidth; i++) {
				t = -1;
				pos = j * (this._mapHeight) + i;
				if (mapFromServer) {
					// console.log("pos:", pos)
					//console.log("got:", mapFromServer[pos])
					if (mapFromServer[pos]) {
						t = mapFromServer[pos]
					}
				} else {
					console.log("no mapFromServer")
				}
				if (t == -1) {
					t = 0;
					// Crafty.math.randomInt(0, 6);
				}
				//console.log("creating", t)
				tile = creationFunc(t);

				tile.addComponent("HexmapNode").hexmapNode(i, j);
				this.placeTile(i, j, tile).setTile(i, j, tile);
			}
		}
	}

	var grass = null;
	var light = null;

	var mapFromServer = [];

	$.ajax({
		url: 'http://localhost:8080/rnd?jsoncallback=?',
		async: true,
		dataType: 'jsonp',
		success: function(data) {

			$.each(data, function(i, v) {
				mapFromServer.push(v);
			});

			hexmap.loadMap(numColumns, numLines, mapFromServer, function(tileIndex) {
				var terrain = tiles[tileIndex];

				// add Mouse back for the MouseOver event.
				return Crafty.e("2D, DOM, Mouse, " + terrain).attr({
					tileIndex: tileIndex
				}).bind("Click", function() {
					var nextTile = (this.tileIndex + 1) % (lastTerrain + 1);
					this.removeComponent(tiles[this.tileIndex])
					this.addComponent(tiles[nextTile])
					this.tileIndex = nextTile;
				}).bind("MouseOver", function() {
					if (grass) {
						if (grass) grass.forEach(function(tile) {
							tile.removeComponent("grass").addComponent("forrest");
						});

						/*
				grass = hexmap.findpath(grass, this);
				if (grass) grass.forEach(function(tile) {
					tile.removeComponent("forrest").addComponent("grass");
				});
				*/
					}

				});
			});


		}
	});
};